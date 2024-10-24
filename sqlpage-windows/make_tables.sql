CREATE TABLE IF NOT EXISTS sue1(
   id INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
  CVE_Number VARCHAR(16) NOT NULL,
  Vulnerability_Score NUMERIC(3,1) NOT NULL,
  Node VARCHAR(100) NOT NULL,
  SW_Make VARCHAR(50) NOT NULL,
  SW_Description VARCHAR(84),
  SW_Version VARCHAR(23) NOT NULL,
  Table_Ref VARCHAR(5) NOT NULL
  );

CREATE TABLE IF NOT EXISTS sue2(
   id INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
  CVE_Number VARCHAR(16) NOT NULL,
  Vulnerability_Score NUMERIC(3,1) NOT NULL,
  Node VARCHAR(100) NOT NULL,
  SW_Make VARCHAR(50) NOT NULL,
  SW_Description VARCHAR(84),
  SW_Version VARCHAR(23) NOT NULL,
  Table_Ref VARCHAR(5) NOT NULL
  );
  
 CREATE TABLE IF NOT EXISTS last_365_days(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );
  
  CREATE TABLE IF NOT EXISTS last_365_days_sue1(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );