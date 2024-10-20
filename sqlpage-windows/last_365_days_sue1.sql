 DROP TABLE IF EXISTS last_365_days_sue1;
 CREATE TABLE IF NOT EXISTS last_365_days_sue1(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );

 INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.29271523178808, 'RedHat Enterprise Linux', 151);
 
 INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.672727272727273, 'FirePower 4125 Next Generation Firewall with Firepower Threat Defense (FTD) Software', 11);
 
 INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.866666666666666, 'Meraki MS425-32 Layer 3 Switch (firmware 2014-09-23)', 9);
 
 INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (7.084210526315789, 'Catalyst 2960-X'), 19;
 
  INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'VirusScan Enterprise', 0);
 
   INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.666666666666667, 'Nessus Vulnerability Scanner', 9);
 
    INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.105263157894737, 'Enterprise Security Information and Event Manager', 19);
 
     INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (7.535509138381201, 'Windows Server 2008 Service Pack 2', 383);
 
      INSERT INTO last_365_days_sue1 (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (8.8, 'Apache OpenOffice', 1);