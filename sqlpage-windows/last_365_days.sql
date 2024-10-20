 DROP TABLE IF EXISTS last_365_days;
 CREATE TABLE IF NOT EXISTS last_365_days(
  Vulnerability_Score FLOAT(15),
  SW_Description VARCHAR(84) NOT NULL,
  Number_Of_Occurrences INT NOT NULL
  );

 INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.29271523178808, 'RedHat Enterprise Linux', 151);
 
 INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.672727272727273, 'FirePower 4125 Next Generation Firewall with Firepower Threat Defense (FTD) Software', 11);
 
 INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.866666666666666, 'Meraki MS425-32 Layer 3 Switch (firmware 2014-09-23)', 9);
 
 INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (7.084210526315789, 'Catalyst 2960-X'), 19;
 
  INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'VirusScan Enterprise', 0);
 
   INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.666666666666667, 'Nessus Vulnerability Scanner', 9);
 
    INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.105263157894737, 'Enterprise Security Information and Event Manager', 19);
 
     INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (7.535509138381201, 'Windows Server 2008 Service Pack 2', 383);
 
      INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (8.8, 'Apache OpenOffice', 1);
 
       INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (6.105504587155964, 'Gitlab Enterprise Edition', 109);
 
        INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'IIS', 0);
 
         INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (8.48, 'VPN Server', 5);
 
          INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (8.425, 'CI/CD pipeline', 4);
 
           INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'August Connect Wi-Fi Bridge phone application', 0);
 
            INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'August Connect device firmware', 0);
 
             INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (NULL, 'Netgear WAC510 Wireless Access Point Firmware', 0);
 
             INSERT INTO last_365_days (Vulnerability_Score, SW_Description, Number_Of_Occurrences) 
 VALUES (8.568571428571428, 'Microsoft SQL Server', 35);