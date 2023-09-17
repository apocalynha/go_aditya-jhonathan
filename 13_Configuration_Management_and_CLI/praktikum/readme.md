# Clean Code
## Install Oh My Zsh!

![omz](/13_Configuration_Management_and_CLI/screenshots/omz.JPG)

## Script Automate

```
export URL="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
export maindir="$1 at $(date)"

mkdir "$maindir"
mkdir "$maindir/about_me"
mkdir "$maindir/about_me/personal"
mkdir "$maindir/about_me/professional"
mkdir "$maindir/my_friends"
mkdir "$maindir/my_system_info"

echo "$2" > "$maindir/about_me/personal/facebook.txt"
echo "$3" > "$maindir/about_me/professional/linkedin.txt"
curl "$URL" -o "$maindir/my_friends/list_of_my_friends.txt"
uname -a > "$maindir/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$maindir/my_system_info/internet_connection.txt"

tree "$maindir"
```

![automate](/13_Configuration_Management_and_CLI/screenshots/automate.JPG)