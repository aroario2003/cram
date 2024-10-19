import pymysql
import pandas as pd

# Load the CSV file from the correct path (use the correct WSL path)
csv_file_path = r"C:\Users\bluel\OneDrive - University of Mary Washington\cves2\sue1\sue1_raw.csv"
csv_data = pd.read_csv(csv_file_path)

# Connect to the MySQL database
connection = pymysql.connect(
    host='localhost',  # if running MySQL locally on Ubuntu
    user='root',       # MySQL root username
    password='abcd',  
    database='alejandro'  # replace with your database name
)

cursor = connection.cursor()

# Prepare the INSERT INTO query
insert_query = """
INSERT INTO sue1 (CVE_Number, Vulnerability_Score, Node, SW_Make, SW_Description, SW_Version, Table_Ref)
VALUES (%s, %s, %s, %s, %s, %s, %s)
"""

# Iterate over the DataFrame and insert the data into the table
for i, row in csv_data.iterrows():
    cursor.execute(insert_query, (
        row['CVE_Number'],
        row['Vulnerability_Score'],
        row['Node'],
        row['SW_Make'],
        row['SW_Description'],
        row['SW_Version'],
        row['Table']
    ))

# Commit the transaction
connection.commit()

# Close the connection
cursor.close()
connection.close()

print("Data inserted successfully!")
