provider "aws"{
    region = "ap-south-1"
}

//create DynamoDB
resource "aws_dynamodb_table" "movies" {
    name = "Movies"
    billing_mode = "PAY_PER_REQUEST"
    hash_key = "movieId"

    attribute {
        name = "movieId"
        type = "S"
    }

    attribute {
        name = "title"
        type = "S"
    }

    attribute {
        name = "genre"
        type = "S"
    }

    global_secondary_index {
        name = "TitleIndex"
        hash_key = "title"
        projection_type = "ALL"
    }

    global_secondary_index {
        name = "genreIndex"
        hash_key = "genre"
        projection_type = "ALL"
    }
}

//create IAM role
resource "aws_iam_role" "lambda_role"{
    name = "movie_lambda_role"

    assume_role_policy = jsonencode({
        Version = "2012-10-17",
        Statement = [{
            Effect = "Allow",
            Principal = {
                Service = "lambda.amazonaws.com"
            },
            Action = "sts:AssumeRole"
        }]
    })
}

//create IAM policies
resource "aws_iam_role_policy" "lambda_policy" {
    role = aws_iam_role.lambda_role.id
    policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Effect = "Allow"
                Action = [
                    "dynamodb:*"
                ],
                Resource = "*"
            },
            {
                Effect = "Allow",
                Action = [
                    "logs:*"
                ],
                Resource = "*"
            }
        ]
    })
}

//create Lambda function
resource "aws_lambda_function" "movie_api"{
    function_name = "movie-api"
    role = aws_iam_role.lambda_role.arn
    handler = "bootstrap"
    runtime = "provided.al2"
    filename = "../build/movie-api.zip"
    source_code_hash = filebase64sha256("../build/movie-api.zip")
    
    environment {
        variables = {
            TABLE_NAME = aws_dynamodb_table.movies.name
        }
    }
}

//create API Gateway
resource "aws_apigatewayv2_api" "http_api"{
    name = "movie-http-api"
    protocol_type = "HTTP"
}
resource "aws_apigatewayv2_integration" "lambda"{
    api_id = aws_apigatewayv2_api.http_api.id
    integration_type = "AWS_PROXY"
    integration_uri = aws_lambda_function.movie_api.invoke_arn
}
resource "aws_apigatewayv2_route" "any"{
    api_id = aws_apigatewayv2_api.http_api.id
    route_key = "ANY /{proxy+}"
    target = "integrations/${aws_apigatewayv2_integration.lambda.id}"
}
resource "aws_apigatewayv2_route" "root"{
    api_id = aws_apigatewayv2_api.http_api.id
    route_key = "ANY /"
    target = "integrations/${aws_apigatewayv2_integration.lambda.id}"
}
resource "aws_apigatewayv2_stage" "default" {
    api_id = aws_apigatewayv2_api.http_api.id
    name = "$default"
    auto_deploy = true
}

//allow API gateway to call lamda
resource "aws_lambda_permission" "api" {
    statement_id = "AllowAPIGateway"
    action = "lambda:InvokeFunction"
    function_name = aws_lambda_function.movie_api.function_name
    principal = "apigateway.amazonaws.com"
    source_arn= "${aws_apigatewayv2_api.http_api.execution_arn}/*/*"
}

output "api_url" {
    value = aws_apigatewayv2_api.http_api.api_endpoint
}