YEAR=$1
DAY=$2

DATA_FILE_NAME="input.txt"

if [ $DEMO ]; then
    DATA_FILE_NAME="input_demo.txt"
fi

sqlite3 "" \
    "create table input(line TEXT);" \
    ".import ./data/$YEAR/$DAY/$DATA_FILE_NAME input" \
    "create table output (question text, answer text);" \
    ".read ./sql/$YEAR/$DAY/script.sql" \
    "select * from output;" \
    ".exit"

