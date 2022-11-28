INSERT INTO
    `dmpr`.`prescriptions` (
        `Id`,
        `MedicineName`,
        `IsActive`,
        `TimesInPeriod`,
        `PeriodLengthInMinutes`,
        `TotalDurationInMinutes`,
        `StartDate`,
        `CountTaken`,
        `CountLeft`
    )
VALUES
    (
        null,
        'Vitamin C',
        0,
        2,        /* 2 times */
        1440,     /* per day */
        10080,    /* for 7 days */
        null,
        0,
        14
    );