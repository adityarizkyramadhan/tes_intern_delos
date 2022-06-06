# tes_intern_delos
# <strong>Overview</strong><br>

- Language ```Go``` <br>
- Framework ```Gin``` and ```Gorm``` <br>
- Database ```Postgres``` with online database e.g. Supabase and ```MySql``` but I run that without docker because my MySql run with Apache (XAMPP) <br>

# <strong>Architecture</strong><br>

- Infrastructure <br>
    => - App = To get all drivers that's needed to run the program, example : jwt secret key in .env file <br>
    => - Database = To get all database needs in .env file<br>
- Domain <br>
    => Entity data <br>
- Implementation <br>
    => Service to interact with database <br>
    => Handler to interact with domain<br>
    => Router to interact with user<br>
- Middleware <br>
    => Authentication <br>
    => Authorization to check if the user is allowed to do the action<br>
- Route <br>
    => Route to get the user to the right page<br>
- Utils <br>
    => Helper function to make response<br>

# <strong>How to run the app</strong><br>
- Install the dependencies <br>
- Docker compose <br>
    => ```I don't know how to run docker compose with mySql but my solution is use supabase because supabase doesnt need docker-compose``` <br>
- Run the app <br>
    => ```go run main.go``` <br>






