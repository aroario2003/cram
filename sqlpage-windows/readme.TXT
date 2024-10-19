i use make_tables.sql first in mysql to create tables
then i change import_csv.py filepath and table name. table name is sue1 or sue2 and filepath depends on you.
i run import_csv.py once for each table, so you have to change the filepath and table name twice
after you do that, http://localhost:8080/form_with_table.sql should work showing you SUE1 tables. until delete cve number is complete, SUE2 won't be made, but would not be hard to make (can copy paste to a different sql file)

hope you can fix delete. from my understanding the only way to make it work would be to give a dropdown menu.
https://sql.datapage.app/component.sql?component=form
https://learnsqlpage.com/
https://github.com/sqlpage/SQLPage/tree/main/examples
https://github.com/sqlpage/SQLPage/discussions