# Split string
This activity let you read a db var in Redis


## Installation

```bash
flogo install github.com/ultrafab/readredisvar_flogo
```
Link for flogo web:
```
https://github.com/ultrafab/readredisvar_flogo
```

## Schema
Inputs and Outputs:

```json
{
 "input":[
     {
       "name": "host",
       "type": "string"
     },
     {
       "name": "port",
       "type": "string"
     },
     {
       "name": "dbNo",
       "type": "integer"
     },
     {
       "name": "variable",
       "type": "string"
     }
  ],
  "output": [
    {
      "name": "value",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| host    | the Redis address |
| port    | the Redis address |
| dbNo    | the Redis database number |
| variable    | the key of the variable to read |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| value    | the value of the variable in Redis |
