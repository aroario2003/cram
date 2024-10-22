CREATE TABLE CVE_List_SUE2 (
    CVE_Number VARCHAR(50),
    Software VARCHAR(255),
    Vulnerability_Score DECIMAL(3,1),
    Time_to_Fix INT,
    Solved BOOLEAN
);
INSERT INTO CVE_List_SUE2 (CVE_Number, Software, Vulnerability_Score, Time_to_Fix, Solved)
VALUES
('CVE-2023-20269', 'Cisco FirePower 4125', 9.1, 5, 0),
('CVE-2023-20256', 'Cisco FirePower 4125', 5.8, 5, 0), 
('CVE-2023-20247', 'Cisco FirePower 4125', 4.3, 5, 0),
('CVE-2023-20200', 'Cisco FirePower 4125', 6.3, 5, 0),
('CVE-2023-20995', 'Cisco FirePower 4125', 8.6, 5, 0),
('CVE-2023-20015', 'Cisco FirePower 4125', 6.7, 5, 0),
('CVE-2023-20934', 'Cisco FirePower 4125', 7.8, 5, 0),
('CVE-2018-0284', 'Cisco Meraki MS425-32', 6.5, 5, 0),
('CVE-2014-7999', 'Cisco Meraki MS425-32', 7.7, 5, 0),
('CVE-2014-7993', 'Cisco Meraki MS425-32', 3.3, 5, 0),
('CVE-2014-7994', 'Cisco Meraki MS425-32', 5.4, 5, 0),
('CVE-2014-7995', 'Cisco Meraki MS425-32', 7.2, 5, 0),
('CVE-2021-4218', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2020-4125', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2020-14312', 'Red Hat Enterprise Linux', 5.9, 0, 0),
('CVE-2020-10730', 'Red Hat Enterprise Linux', 6.5, 0, 0),
('CVE-1999-0894', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2000-1207', 'Red Hat Enterprise Linux', 7.2, 0, 0),
('CVE-2004-0217', 'Red Hat Enterprise Linux', 7.0, 0, 0),
('CVE-2007-0980', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2007-2797', 'Red Hat Enterprise Linux', 1.6, 0, 0),
('CVE-2008-0889', 'Red Hat Enterprise Linux', 1.6, 0, 0),
('CVE-2011-0536', 'Red Hat Enterprise Linux', 6.9, 0, 0),
('CVE-2013-5364', 'Red Hat Enterprise Linux', 3.6, 0, 0),
('CVE-2017-8989', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2018-1113', 'Red Hat Enterprise Linux', 5.3, 0, 0),
('CVE-2020-10759', 'Red Hat Enterprise Linux', 6.0, 0, 0),
('CVE-2021-20269', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2022-0851', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2022-0852', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2022-28623', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2003-0019', 'Red Hat Enterprise Linux', 7.2, 0, 0),
('CVE-2003-0080', 'Red Hat Enterprise Linux', 9.4, 0, 0),
('CVE-2019-10214', 'Red Hat Enterprise Linux', 5.9, 0, 0),
('CVE-2019-19339', 'Red Hat Enterprise Linux', 6.5, 0, 0),
('CVE-2020-14306', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2020-14391', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2020-15719', 'Red Hat Enterprise Linux', 4.2, 0, 0),
('CVE-2020-1702', 'Red Hat Enterprise Linux', 3.3, 0, 0),
('CVE-2021-20325', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2021-43816', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2022-0552', 'Red Hat Enterprise Linux', 5.9, 0, 0),
('CVE-2022-1665', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2023-2203', 'Red Hat Enterprise Linux', 10.0, 0, 0),
('CVE-2023-2295', 'Red Hat Enterprise Linux', 9.4, 0, 0),
('CVE-2023-4042', 'Red Hat Enterprise Linux', 5.5, 0, 0),
('CVE-2012-0158', 'Microsoft SQL Server 2008 SP2', 8.8, 5, 0),
('CVE-2015-1762', 'Microsoft SQL Server 2008 SP2', 7.1, 5, 0),
('CVE-2015-1763', 'Microsoft SQL Server 2008 SP2', 8.5, 5, 0),
('CVE-2012-1856', 'Microsoft SQL Server 2008 SP2', 8.8, 5, 0),
('CVE-2023-47804', 'Apache OpenOffice', 8.8, 5, 0),
('CVE-2022-37401', 'Apache OpenOffice', 8.8, 5, 0),
('CVE-2021-33035', 'Apache OpenOffice', 7.8, 5, 0),
('CVE-2020-13958', 'Apache OpenOffice', 7.8, 5, 0),
('CVE-2017-12607', 'Apache OpenOffice', 7.8, 5, 0),
('CVE-2019-17098', 'August Wi-Fi Bridge', 3.1, 5, 0),
('CVE-2019-17518', 'August Connect Device Firmware', 6.5, 5, 0),
('CVE-2018-21133', 'Netgear WAC510 Firmware', 9.8, 5, 0),
('CVE-2018-21132', 'Netgear WAC510 Firmware', 9.8, 5, 0),
('CVE-2018-21131', 'Netgear WAC510 Firmware', 9.1, 5, 0),
('CVE-2018-21130', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2018-21129', 'Netgear WAC510 Firmware', 6.5, 5, 0),
('CVE-2018-21128', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2018-21127', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2018-21126', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2018-21125', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2018-21124', 'Netgear WAC510 Firmware', 8.8, 5, 0),
('CVE-2016-6473', 'Cisco Catalyst 2960-X', 6.5, 5, 0),
('CVE-2017-6606', 'Cisco Catalyst 2960-X', 6.4, 5, 0),
('CVE-2017-3803', 'Cisco Catalyst 2960-X', 4.7, 5, 0),
('CVE-2016-1425', 'Cisco Catalyst 2960-X', 6.5, 5, 0),
('CVE-2024-1495', 'GitLab Enterprise Edition', 6.5, 5, 0),
('CVE-2024-2743', 'GitLab Enterprise Edition', 9.1, 5, 0),
('CVE-2024-2576', 'GitLab Enterprise Edition', 4.3, 5, 0),
('CVE-2010-1256', 'Microsoft IIS', 8.5, 5, 0),
('CVE-2010-3972', 'Microsoft IIS', 10.0, 5, 0),
('CVE-2022-0547', 'OpenVPN Server', 9.8, 5, 0),
('CVE-2020-15078', 'OpenVPN Server', 7.5, 5, 0),
('CVE-2020-20813', 'OpenVPN Server', 7.5, 5, 0),
('CVE-2017-12166', 'OpenVPN Server', 9.8, 5, 0),
('CVE-2017-1000353', 'Jenkins CI/CD', 9.8, 5, 0),
('CVE-2017-1000354', 'Jenkins CI/CD', 8.8, 5, 0),
('CVE-2017-1000356', 'Jenkins CI/CD', 8.8, 5, 0),
('CVE-2015-7833', 'RedHat RHEL 7.1', 4.9, 5, 0),
('CVE-2020-7337', 'McAfee VirusScan Enterprise', 6.7, 5, 0),
('CVE-2009-5118', 'McAfee VirusScan Enterprise', 9.3, 5, 0),
('CVE-2007-2152', 'McAfee VirusScan Enterprise', 7.9, 5, 0),
('CVE-2023-0101', 'Tenable Nessus', 8.8, 5, 0),
('CVE-2021-20135', 'Tenable Nessus', 6.7, 5, 0),
('CVE-2020-5765', 'Tenable Nessus', 5.4, 5, 0),
('CVE-2024-23675', 'Splunk SIEM', 6.5, 5, 0),
('CVE-2024-23676', 'Splunk SIEM', 3.5, 5, 1),
('CVE-2023-40593', 'Splunk SIEM', 7.5, 5, 0),
('CVE-2023-40592', 'Splunk SIEM', 6.1, 5, 0),
('CVE-2017-8543', 'Microsoft Windows Server 2008 SP2', 9.8, 5, 0),
('CVE-2014-0301', 'Microsoft Windows Server 2008 SP2', 9.3, 5, 0),
('CVE-2014-0323', 'Microsoft Windows Server 2008 SP2', 6.6, 5, 0),
('CVE-2014-0315', 'Microsoft Windows Server 2008 SP2', 6.9, 5, 0),
('CVE-2013-5058', 'Microsoft Windows Server 2008 SP2', 6.9, 5, 0),
('CVE-2013-5056', 'Microsoft Windows Server 2008 SP2', 9.3, 5, 0);
