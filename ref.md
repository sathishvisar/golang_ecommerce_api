Successful Responses (2xx):

200 OK: The request was successful.

Message: "Product found successfully."
201 Created: The request resulted in a new resource being created.

Message: "Product created successfully."
204 No Content: The request was successful, and there is no additional information to send back.

Message: (No message body)
Client Error Responses (4xx):

400 Bad Request: The request could not be understood or was missing required parameters.

Message: "Bad request. Please check your request parameters."
401 Unauthorized: Authentication failed or user lacks necessary permissions.

Message: "Unauthorized. Please authenticate and try again."
403 Forbidden: The server understood the request but refuses to authorize it.

Message: "Forbidden. You don't have permission to access this resource."
404 Not Found: The requested resource could not be found.

Message: "Product not found."
422 Unprocessable Entity: The request was well-formed but unable to be followed due to semantic errors.

Message: "Unprocessable entity. Please check your request data."
Server Error Responses (5xx):

500 Internal Server Error: An unexpected condition was encountered by the server.