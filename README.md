# linkly

A simple url shortner.

## Endpoints

| Methods | Endpoints             | Description                       |
| ------- | --------------------- | --------------------------------- |
| GET     | /api/health           | health checking                   |
| GET     | /:short_url           | url redirection                   |
| GET     | /api/links            | Get all short urls                |
| GET     | /api/links/:short_url | Get information about a short url |
| POST    | /api/links            | Create short url                  |
| PATCH   | /api/links            | Update short url                  |
| DELETE  | /api/links            | Delete short url                  |

## Author

- Injamul Mohammad Mollah <mrinjamul@gmail.com>
