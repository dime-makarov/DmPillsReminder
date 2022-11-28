CREATE DATABASE `dmpr`
/*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */
/*!80016 DEFAULT ENCRYPTION='N' */
;

CREATE TABLE `dmpr`.`prescriptions` (
    `Id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `MedicineName` NVARCHAR(100) NOT NULL,
    `IsActive` BIT NULL DEFAULT 0,
    `TimesInPeriod` INT UNSIGNED NULL DEFAULT 0,
    `PeriodLengthInMinutes` INT UNSIGNED NULL DEFAULT 0,
    `TotalDurationInMinutes` INT UNSIGNED NULL DEFAULT 0,
    `StartDate` DATETIME NULL,
    `CountTaken` INT UNSIGNED NULL DEFAULT 0,
    `CountLeft` INT UNSIGNED NULL DEFAULT 0,
    PRIMARY KEY (`Id`)
);