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

# The 2nd input is the path to the folder, if not given, use "."
path=$2

if [ -z "$path" ]; then
    path = "."
fi

# Ensure folder does not exist:
if [ -d "$path/day$day" ]; then
    echo "Folder day$day already exists, aborting..."
    exit 1
fi

# Ensure dayx folder exists:
if ! [ -d "$path/dayx" ]; then
    echo "Folder dayx does not exist, aborting..."
    exit 1
fi

cp -r $path/dayx $path/day$day

# Go into the file main.go.stub and change "0," to "$day,", but without the leading 0
if [ "$(uname)" == "Darwin" ]; then
    sed -i '' "s/0,/$(echo $day | sed 's/^0*//'),/g" $path/day$day/main.go.stub
else
    sed -i "s/0,/$(echo $day | sed 's/^0*//'),/g" $path/day$day/main.go.stub
fi

#  Then change the name of the file main.go.stub to main.go
mv $path/day$day/main.go.stub $path/day$day/main.go

# Done!
echo "Done!"
