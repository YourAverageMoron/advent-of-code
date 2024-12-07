with a as (SELECT
    ROW_NUMBER () OVER ( 
        ORDER BY cast(SUBSTR(line, 1, INSTR(line, '   ') - 1) as INTEGER)
    ) id,
    cast(SUBSTR(line, 1, INSTR(line, '   ') - 1) as INTEGER) AS value
FROM INPUT),

b as (SELECT
    ROW_NUMBER () OVER ( 
        ORDER BY  cast(SUBSTR(line, INSTR(line, '   ') + 1) as INTEGER)
    ) id,
    cast(SUBSTR(line, INSTR(line, '   ') + 1) as INTEGER) AS value
FROM INPUT)


insert into output (question, answer) values (
    'Question 1',
    (select sum(abs(a.value - b.value)) from a join b on a.id = b.id)
);

------------------

with a as (SELECT
    cast(SUBSTR(line, 1, INSTR(line, '   ') - 1) as INTEGER) AS value
FROM INPUT),

b as (SELECT
    cast(SUBSTR(line, INSTR(line, '   ') + 1) as INTEGER) AS value
FROM INPUT),

c as (select value, count(value) c from b group by value)

insert into output (question, answer) values (
    'Question 2',
    (select sum(a.value * c.c) from a join c on a.value = c.value)
);
