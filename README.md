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

## Get All Cars
GET /cars
Retrieves a list of all cars in the system.

## Get Car by ID
GET /cars/{id}
Retrieves information about a specific car identified by its ID.

## Update Car Status
PUT /cars/{id}
Updates the status of a specific car identified by its ID.

## Remove Car
DELETE /cars/{id}
Removes a car from the system when service is done.

## Author
Shresth Agarwal
Student ID: 20102116
Email: Shresth984@gmail.com
