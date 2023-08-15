create table transactions
(
    id               uuid             not null,
    id_from          integer          not null,
    id_to            integer          not null,
    amount           double precision not null,
    currency_type    varchar          not null,
    transaction_type varchar          not null
);