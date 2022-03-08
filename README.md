# Todo Api
Example application for api introduction

## Endpoints
| Path                            | Verb   | Description                                |
| ---                             | ---    | ---                                        |
| /lists                          | GET    | Retrieve a list of TODO lists              |
| /lists                          | POST   | Create a new TODO list                     |
| /lists/:slug                    | PATCH  | Update a TODO list                         |
| /lists/:slug                    | DELETE | Delete a TODO list                         |
| /lists/:slug/entries            | GET    | Get a list of the entries on the TODO list |
| /lists/:slug/entries            | POST   | Create a new entry on the todo list        |
| /lists/:slug/entries/:id        | PATCH  | Update an entry on the TODO list           |
| /lists/:slug/entries/:id/toggle | POST   | Toggle the done state on the entry         |
| /lists/:slug/entries/:id        | DELETE | Remove an entry on the TODO list           |

## JSON Output
All responses will contain both the `outcome` and `message` keys in the response json,

On successful requests the server will return a 200 status code `outcome` will be true and the `message` key will be null.

On unsuccessful responses the status code will be a veriety of codes depending on the error, `outcome` will be false
and the `message` will be set

For any requist with additional output i have given an example below

## Endpoint url's
Some of the urls contain either `:slug` or `:id`
there are to be replaced by the lists slug and entries id respectively

### /lists (GET - fetch a list of lists)
This endpoint takes no parameters
```json
{
    "outcome": true,
    "message": null,
    "lists": [
        {
            "id": "[int]",
            "created_at": "[string]",
            "updated_at": "[string]",
            "slug": "[string]",
            "title": "[string]"
        }
    ]
}
```

### /lists (POST - create a new list)
| Key | Type | Description |
| --- | --- | --- |
| slug | string | A unique key for the list (this is what would traditionally be used in the url) |
| title | string | Display title for the list |

### /lists/:slug (PATCH - Update a list)
| Key | Type | Description |
| --- | --- | --- |
| title | string | Display title for the list |

### /lists/:slug (DELETE - Delete a list)
This endpoint takes no parameters

### /lists/:slug/entries (GET - List all the todo entries in a list)
This endpoint takes no parameters

```json
{
    "outcome": true,
    "message": null,
    "list": {
        "id": "[int]",
        "created_at": "[string]",
        "updated_at": "[string]",
        "slug": "[string]",
        "title": "[string]"
    },
    "entries": [
        {
            "id": "[int]",
            "created_at": "[string]",
            "updated_at": "[string]",
            "text": "[string]",
            "done": "[bool]"
        }
    ]
}
```

### /lists/:slug/entries (POST - create a new list entry)
| Key | Type | Description |
| --- | --- | --- |
| text | string | The entry text |

### /lists/:slug/entries/:id (PATCH - update an existing list entry)
| Key | Type | Description |
| --- | --- | --- |
| text | string | The entry text |

### /lists/:slug/entries/:id (DELETE - remove a list entry)
This endpoint takes no parameters

### /lists/:slug/entries/:id/toggle (POST - Toggle the completion state of the list)
This endpoint takes no parameters

