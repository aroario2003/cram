CREATE TABLE CVE_List_SUE1 (
    CVE_Number VARCHAR(50) PRIMARY KEY,
    Software VARCHAR(255),
    Vulnerability_Score DECIMAL(3,1),
    Time_to_Fix INT,
    Solved BOOLEAN
);

INSERT INTO CVE_List_SUE1 (CVE_Number, Software, Vulnerability_Score, Time_to_Fix, Solved) VALUES
('CVE-2023-20269', 'Cisco Firepower', 9.1, 0, 0),
('CVE-2023-20256', 'Cisco Firepower', 5.8, 0, 0),
('CVE-2023-20247', 'Cisco Firepower', 4.3, 0, 0),
('CVE-2018-0284', 'Cisco Meraki', 6.5, 0, 0),
('CVE-2014-7999', 'Cisco Meraki', 7.7, 0, 0),
('CVE-2016-6473', 'Cisco Catalyst', 6.5, 0, 0),
('CVE-2020-7337', 'McAfee VirusScan', 6.7, 0, 0),
('CVE-2012-2697', 'RHEL5', 4.9, 0, 0),
('CVE-2024-23675', 'Tenable Nessus', 6.5, 0, 0),
('CVE-2023-47804', 'Apache OpenOffice', 8.8, 0, 0);
