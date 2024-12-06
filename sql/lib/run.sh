YEAR=$1
DAY=$2

sqlite3 "" \
    "create table input(line TEXT);" \
    ".import ./data/$YEAR/$DAY/input.txt input" \
    ".read ./sql/$YEAR/$DAY/script.sql" \
    "select * from output;" \
    ".exit"

