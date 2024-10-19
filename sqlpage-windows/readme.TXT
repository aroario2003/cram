first you have to install sqlpage. then u can set up tables locally create them and fill with data. then use sqlpage.exe on windows, something else for linux/mac. this has only been tested on windows 10 from my understanding. sqlpage in general is used to display databases

i use make_tables.sql first in mysql to create tables
then i change import_csv.py filepath and table names. table name is sue1 or sue2 and filepath depends on your system, but are in the csvs folder of this directory
i run import_csv.py once for each table, so you have to change the filepath and table name twice
also use localhost:8080 to access website and change sqlpage.json in sqlpage folder to match your database login credentials
then run sqlpage.exe windows or whatever os u have version.
after you do that, http://localhost:8080/form_with_table.sql should work showing you SUE1 tables. until delete cve number is complete, SUE2 won't be made, but would not be hard to make (can copy paste to a different sql file)

we were using some links like this to help with coding:
https://sql.datapage.app/component.sql?component=form
https://learnsqlpage.com/
https://github.com/sqlpage/SQLPage/tree/main/examples
https://github.com/sqlpage/SQLPage/discussions
