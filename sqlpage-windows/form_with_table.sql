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
  
 CREATE TABLE IF NOT EXISTS last_365_days_sue1(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );
  
 CREATE TABLE IF NOT EXISTS last_365_days_sue1_sue1(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );
  
  create table if not exists csv2sql_4758 (
	id int not null auto_increment primary key, 
	cveid varchar(64),
	vendorproject varchar(64),
	product text,
	vulnerabilityname text,
	dateadded varchar(64),
	shortdescription text,
	requiredaction text,
	duedate varchar(64),
	knownransomwarecampaignuse varchar(64),
	notes text,
	cwes varchar(64)
	
);
  
 SELECT 'list' AS component, 'You\'re in the SUE1 page' AS title;
 SELECT 'Home' AS title, 'NSWCDD Hackathon Hub' AS description,  '/' AS link;



-- Delete the record
-- doesn't even work on it's own
--DELETE FROM sue1 WHERE CVE_Number=$del;
DELETE FROM sue1 WHERE id=$delete;
DELETE FROM sue1 WHERE CVE_Number = $del collate utf8mb4_0900_ai_ci


--SELECT 'steps' as component;
SELECT 'title' as component,
--AS contents;
CONCAT('Security Score: ', .6*
(100-(
SELECT LEAST((
SELECT GREATEST((
SELECT IFNULL(SUM(
POW(
	(
	SELECT 
		LEAST(
			CASE 
				WHEN Vulnerability_Score < 2.5 THEN Vulnerability_Score * 0.75
				WHEN Vulnerability_Score > 7.5 THEN Vulnerability_Score * 1.25
				ELSE Vulnerability_Score
			END, 
		10) AS adjusted_vulnerability_score_limited
	)
,2)) 
/ COUNT(*), "0") FROM sue1)
*$criticality
,0))
,100)
))

+
(.2*(100-(
SELECT LEAST((
SELECT GREATEST((
SELECT IFNULL(
SUM(POW
(

(	SELECT 
		LEAST(
			CASE 
				WHEN last_365_days_sue1.Vulnerability_Score < 2.5 THEN last_365_days_sue1.Vulnerability_Score * 0.75
				WHEN last_365_days_sue1.Vulnerability_Score > 7.5 THEN last_365_days_sue1.Vulnerability_Score * 1.25
				ELSE last_365_days_sue1.Vulnerability_Score
			END, 
		10) AS adjusted_vulnerability_score_limited
	),2)
)
/COUNT(sue1.SW_Description)
,(select pow(
sum(vulnerability_score
*number_of_occurrences)
 / sum(number_of_occurrences),2) from last_365_days_sue1))
 FROM last_365_days_sue1 
 INNER JOIN sue1 ON last_365_days_sue1.SW_Description=sue1.SW_Description)*
 $criticality
 ,0))
,100))))

 -(.2*10*(
 SELECT LEAST((
SELECT GREATEST((
SELECT IFNULL((
 SELECT 
    MAX(sub.Vulnerability_Score) AS Max_Vulnerability_Score
FROM 
    (SELECT 
        SUE1.Vulnerability_Score 
     FROM 
        csv2sql_4758 
     INNER JOIN 
        SUE1 ON SUE1.CVE_Number = csv2sql_4758.cveid) sub
)*$criticality
 ,0))
,10))
, "-10")
 ))


, ' / 100') AS contents;

select 
    'title'   as component,
    'Click a criticality button to calculate the security score' as contents where $criticality is null;

SELECT
	'button' as component;
SELECT
	'/insert_into_sue1.sql' as link,
CONCAT((SELECT IF((SELECT COUNT(TABLE_NAME) FROM information_schema.TABLES WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME='SUE1') = 0, "Create", "Recreate")), ' SUE1 Table')  as title;
SELECT
	'?criticality=1.25' as link,
	'danger' as color,
	'High Criticality' as title;
SELECT
	'?criticality=1' as link,
	'warning' as color,
	'Medium Criticality' as title;
SELECT
	'?criticality=.75' as link,
	'success' as color,
	'Low Criticality' as title;

-- Display the table with actions
SELECT 'table' AS component,
    'CVE' AS markdown, 'ID2' AS markdown, TRUE AS sort, TRUE AS search;
SELECT  
	 id AS ID,
     CVE_Number AS "CVE Number", 
     Vulnerability_Score AS "Vulnerability Score", 
     Node AS Node,
	 SW_Make AS "Software Make",
	 SW_Description AS "Software Description",
	 SW_Version AS "Software Version",
	 Table_REF as "Table",
    '[🗑️CVE](?del=' || CVE_Number || ')' AS CVE,          -- Dynamic link for edit
    '[🗑️ID](?delete=' || id || ')' AS ID2        -- Dynamic link for delete
FROM sue1;