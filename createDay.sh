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
if [ "$(uname)" == "Darwin" ]; then
sed -i '' "s/dayx/day$day/g" day$day/main.js.stub
sed -i '' "s/day: ''/day: $day/g" day$day/main.js.stub
else
sed -i "s/dayx/day$day/g" day$day/main.js.stub
sed -i "s/day: ''/day: $day/g" day$day/main.js.stub
fi

#  Then change the name of the file main.js.stub to main.js
mv day$day/main.js.stub day$day/main.js

# Go into ./main.js, and add a new line "import * as dayx from "./dayx/main.js";" on the first empty line, then a gap after
if [ "$(uname)" == "Darwin" ]; then
sed -i '' "/^$/i\\
import * as day$day from \"./day$day/main.js\";\\
" ./main.js
else
sed "/^$/i\\
import * as day$day from \"./day$day/main.js\";\\
" ./main.js
fi

# Now add 2 new lines to the end of the file, 'dayx.partOne();' and 'dayx.partTwo();'
echo "day$day.partOne();" >>main.js
echo "day$day.partTwo();" >>main.js

# Done!
echo "Done!"
