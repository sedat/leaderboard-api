# Leaderboard API

A game leaderboard implemented in Golang, Redis and MongoDB. Includes general and country based ranking with Redis Sorted Sets, MongoDB as persistent database and password hash encryption with JWT authentication. 

## Public url

`https://arcane-beach-01523.herokuapp.com/`

The app is deployed to Heroku using Heroku Redis, and a Mongodb cluster on Mongodb Atlas. All free tier.

## Running locally

To run the application locally you need;

- Golang
- MongoDB
- Redis

All these tools above are configured to run on their local default configurations.

Once you completed setting them up, please run `god mod download` to download Go dependencies. You can see the list of dependencies in `go.mod`. Please don't forget to add the required environment variables.

## Running with docker

If you prefer you can run the application in a docker container by typing;

`docker-compose -f docker-compose.yml up`

# Endpoints

## Leaderboard API

Various endpoint that return leaderboard in different ways. All endpoints under Leaderboard requires a token to be in the Authorization header

#### GET /leaderboard

Returns the entire leaderboard of all users. This request takes a lot of time. Heroku as a timeout of 30 seconds. In large collections of data this will return a 503 - Request Timeout. It's better to use paging with the range endpoint considering both time and response payload.

##### Response

```
[
  {
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "0ioYXaaaaaa3V",
    "country": "NZ",
    "points": 10000,
    "rank": 1,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a3bb",
    "display_name": "DOMsOa",
    "country": "VE",
    "points": 10000,
    "rank": 2,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18418",
    "display_name": "UOMNO",
    "country": "NI",
    "points": 10000,
    "rank": 3,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18d37",
    "display_name": "5gHkycaaaaM2rRF",
    "country": "VI",
    "points": 9999,
    "rank": 4,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a2b9",
    "display_name": "kMiJobaaaayAucg",
    "country": "MG",
    "points": 9998,
    "rank": 5,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a236",
    "display_name": "o16LIdaaaae",
    "country": "UA",
    "points": 9998,
    "rank": 6,
    "last_score_timestamp": ""
  }
]
```

#### GET /leaderboard/limit/:limit

Returns the leaderboard up to the number e.g. 4

##### Response

```
[
  {
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "0ioYXaaaaaa3V",
    "country": "NZ",
    "points": 10000,
    "rank": 1,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a3bb",
    "display_name": "DOMsOa",
    "country": "VE",
    "points": 10000,
    "rank": 2,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18418",
    "display_name": "UOMNO",
    "country": "NI",
    "points": 10000,
    "rank": 3,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18d37",
    "display_name": "5gHkycaaaaM2rRF",
    "country": "VI",
    "points": 9999,
    "rank": 4,
    "last_score_timestamp": ""
  },
]
```

#### GET /leaderboard/country/:country

Returns the leaderboard of a country e.g. VI

##### Response

```
[
  {
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "0ioYXaaaaaa3V",
    "country": "VI",
    "points": 10000,
    "rank": 1,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a3bb",
    "display_name": "DOMsOa",
    "country": "VI",
    "points": 9999,
    "rank": 2,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18418",
    "display_name": "UOMNO",
    "country": "VI",
    "points": 9999,
    "rank": 3,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18d37",
    "display_name": "5gHkycaaaaM2rRF",
    "country": "VI",
    "points": 9998,
    "rank": 4,
    "last_score_timestamp": ""
  },
]
```

#### GET /leaderboard/country/:country/limit/:limit

Returns the leaderboard of a country e.g. VI with a limit

##### Response

```
[
  {
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "0ioYXaaaaaa3V",
    "country": "VI",
    "points": 10000,
    "rank": 1,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a3bb",
    "display_name": "DOMsOa",
    "country": "VI",
    "points": 9999,
    "rank": 2,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18418",
    "display_name": "UOMNO",
    "country": "VI",
    "points": 9999,
    "rank": 3,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18d37",
    "display_name": "5gHkycaaaaM2rRF",
    "country": "VI",
    "points": 9998,
    "rank": 4,
    "last_score_timestamp": ""
  },
]
```

#### GET /leaderboard/range/:start/:end

Returns the leaderboard of in a range, can be used for paging. e.g. start: 0, end: 3

##### Response

```
[
  {
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "0ioYXaaaaaa3V",
    "country": "VI",
    "points": 10000,
    "rank": 1,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff1a3bb",
    "display_name": "DOMsOa",
    "country": "VI",
    "points": 9999,
    "rank": 2,
    "last_score_timestamp": ""
  },
  {
    "_id": "5fd6987b6cd5f9a6aff18418",
    "display_name": "UOMNO",
    "country": "VI",
    "points": 9999,
    "rank": 3,
    "last_score_timestamp": ""
  }
]
```

## User API

#### GET /user/fill/:number

Works only on local environment using the [mgodatagen](https://github.com/feliixx/mgodatagen) tool, please install it to use this endpoint.
This endpoint is designed to fill the database with mock data and feed the data into Redis. You have four options for number, 10, 100, 500 and 1000. These numbers all follow three zero's. Which means that 10 is 10000 users, 100 is 100000 users etc. Keep in mind that this operation takes time as Redis is single-threaded.
You should wait till you get the response below.

##### Response

```
{
    "msg": "Database created and redis is initialized with leaderboard"
},

```

#### POST /user/create

Creates a user

##### Body

```
{

    "display_name": "Super Mario",
    "country": "IT",
    "password": "1234" // to be converted and stored as a hash
}
```

##### Response

```
{
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "Super Mario",
    "password" "1234",
    "rank": 100,
    "last_score_timestamp": ""
},

```

#### POST /user/login

User login, returns a token to be used in the Header Authorization key.

##### Body

```
{

    "display_name": "Super Mario",
    "password": "1234"
}
```

##### Response

```
{
  token: <token>
}
```

#### POST /user/score/submit

Add the score to the users point field and updates her rank

##### Body

```
{

    "user_id": "5fd6987b6cd5f9a6aff1a6a8",
    "score_worth": 1234,
    "timestamp": timestamp
}
```

##### Response

```
{
    "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "Super Mario",
    "points": 2345,
    "rank": 50,
    "last_score_timestamp": timestamp
}
```

#### GET /user/profile/:profileId

Returns the user with id profileID. Requires a token to be in the Authorization header

##### Response

```
{
   "_id": "5fd6987b6cd5f9a6aff1a6a8",
    "display_name": "Super Mario",
    "points": 2345,
    "rank": 50,
    "last_score_timestamp": timestamp
}
```
