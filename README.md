# Car Management System

This is a  Go program that implements a basic Car Management System with CRUD operations using RESTful API.

## Table of Contents
- [Introduction](#introduction)
- [Endpoints](#endpoints)
- [Author](#author)

## Introduction

The Car Management System allows you to perform the following operations on a list of cars stored in memory:

- Add a new car
- Retrieve a list of all cars
- Retrieve information about a specific car
- Update the status of a car
- Remove a car from the system when service is done

## Endpoints

## Add Car
POST /cars
- Adds all the cars in Garage 

## Get All Cars
GET /cars
- Retrieves a list of all cars in the system.

## Get Car by ID
GET /cars/{id}
- Retrieves information about a specific car identified by its ID.

## Update Car Status
PUT /cars/{id}
- Updates the status of a specific car identified by its ID.

## Remove Car
DELETE /cars/{id}
- Removes a car from the system when service is done.

## DETAILS
- In the provided Go code, the application is configured to listen on port 5000 for incoming HTTP requests. The line http.ListenAndServe(":5000", router) specifies that the 
  application will be accessible at http://localhost:5000. You can customize the port number according to your preferences or requirements.

- The Postman API collection file for this project is attached in the GitHub repository for easy testing and interaction with the API. You can find it in the postman
  directory.

## Summary
- Add Car: POST /cars - Adds a new car to the system.
- Get All Cars: GET /cars - Retrieves a list of all cars in the system.
- Get Car by ID: GET /cars/{id} - Retrieves information about a specific car identified by its ID.
- Update Car Status: PUT /cars/{id} - Updates the status of a specific car identified by its ID.
- Remove Car: DELETE /cars/{id} - Removes a car from the system when service is done.

## Author
Shresth Agarwal
Student ID: 20102116
Email: Shresth984@gmail.com
