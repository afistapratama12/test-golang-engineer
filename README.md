# test-golang-engineer
for test golang engineer refactory

## LOGIC TEST
- answer about logic test in directory logicTest
- you can read all answer code in logicFunc/logic.go
- or you can running all answer test with ```go run main.go```

## JSON MANIPULATION
- answer about JSON Manipulation test in directory JSONManipulation
- you can read all answer code in main.go
- or you can running JSON Manipulation test with ```go run main.go```

## GOLANG ENGINEER - OAUTH GOOGLE
- answer about golang engineer test in directory oauthGoogle
- you can read all answer code in directory oauthGoogle

- application flow : 
1. you can run the application with the command ```go run main.go```
2. don't forget, before running the application, please prepare the required environment variables by creating an .env file containing important data such as: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, PASS_USER_GENERATE DB_USER, DB_PASS, DB_HOST, DB_NAME. I don't provide my environment variable because it is important data and afraid of misuse. Please understand
3. the application will run on localhost with port 8080 (localhost:8080)
4. you can open the route ```localhost:8080/``` then the application will direct to login using google account at ```localhost:8080/login```
5. if successful the application will redirect to ```localhost:8080/callback``` with json data display. The data sent is the data that was successfully saved to the database from the data that Google provided

- database schema : The database in this application only consists of a user table with columns id, email, password, name, avatar, createdAt and updatedAt. The schema created does not have any relation and the data created is automatically filled in from google login login

- diagram database : you can see the database diagram in the picture : database Diagram.png