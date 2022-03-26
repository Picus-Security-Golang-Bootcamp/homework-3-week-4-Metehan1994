# Homework-3 | Week 4 | Booklist App

## Overview

A prototype of program reading CSV file, connecting database on PostGreSQL and creating two tables on it which are "Author" and "Book" is created.

## How to Use the App ?

The program works with some queries which can be reached from main function.

* List names of book and author names
* Search books and authors with their names and IDs
* Create new books and authors with their info needed
* Delete books and authors from their tables with their names and IDs
* Buy books through book repo and update its quantity in stock
* Get author and book info together from author and book tables and more

### Some Notes for Usage

1. Program produces error messages when it is executed without considering its usage.

2. After deleting a book or an author, its status is changed and kept in the database (soft deleting).

## Package Used

* The program is created with **GO main package & GORM & Godotenv**.
