#bash

if ! [ -z "$1" ]; then
    echo "Day: $1"
    day=$1
else
    echo -n "Day: "
    read day
    echo

    if [ -z "$day" ]; then
        echo "No day given, aborting..."
        exit 1
    fi
fi

# Ensure folder does not exist:
if [ -d "day$day" ]; then
    echo "Folder day$day already exists, aborting..."
    exit 1
fi
cp -r dayx day$day

# Go into the file main.js.stub and change "day: ''" to "day: 'day$day'", and change "dayx" to "day$day"
sed -i '' "s/dayx/day$day/g" day$day/main.js.stub
sed -i '' "s/day: ''/day: $day/g" day$day/main.js.stub

#  Then change the name of the file main.js.stub to main.js
mv day$day/main.js.stub day$day/main.js

# Done!
echo "Done!"
