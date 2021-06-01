# ChatBoard

## How to make test database
```sql
insert into users values(0, gen_random_uuid(), 'user0', 'user0@abc.com', 'user00', now());
insert into users values(1, gen_random_uuid(), 'user1', 'user1@abc.com', 'user01', now());
insert into threads values(0, gen_random_uuid(), 'Topic1 US', 0, now());
insert into posts values(0, gen_random_uuid(), 'The United States of America (U.S.A. or USA), commonly known as the United States (U.S. or US) or America, is a country primarily located in North America. ', 0, 0,now());
insert into posts values(1, gen_random_uuid(), 'It consists of 50 states, a federal district, five major unincorporated territories, 326 Indian reservations, and some minor possessions.', 0, 0,now());
```