# PR Tracker

<p>
`PR Tracker` is an web application made with Golang, Gin, that assists the weight lifting process by logging your achievements and graphically displaying the record.
</p>

## Running to Application Locally via CLI

```console
go run main.go
```

### Register 
```console
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<<USERNAME>>", "password":"<<PASSWORD>>"}' \
    http://localhost:8000/auth/register
```

### Login
```console
curl -i -H "Content-Type: application/json" \
    -X POST \                            
    -d '{"username":<<USERNAME>>, "password":<<PASSWORD>>}' \    http://localhost:8000/auth/login 
```

Once logged in you have access to the following commands:

- **POST** `/workout`
    - Creates a Workout
- **GET** `/workout`
    - Gets all the users workouts
- **DELETE** `/workout/:id`
    - Deletes a users workout given the Workout ID
- **PATCH** `/workout/:id`
    - Updates the users workout information given the ID and the updated parameters
- **POST** `/workout/:id/records`
    - Adds a new record to the workout given the ID of the workout and the new record
- **GET** `/workout/:id/records`
    - Gets all of the records of the workout given the ID of the workout
- **DELETE** `/workout/:id/records/:recordId`
    - Deletes a record of a workout given the ID of the workout and the record ID

### File Structure

The overall file structure is layered such that each folder is responsible for similar actions and can imported by the folder name as the package name.

- `api`
    - `controller`
        - `authentication.go`
        - `record.go`
        - `workout.go`
    - `database`
        - `datase.go`
    - `helper`
        - `jwt.go`
    - `middleware`
        - `jwtAuth.go`
    - `model`
        - `authenticationInput.go`
        - `record.go`
        - `user.go`
        - `workout.go`