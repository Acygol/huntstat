# Summary
HuntStat is a Discord bot for TheHunter: Classic communities to automate their administration.

# Features
1. (Un)register community members

![](https://i.imgur.com/l17SOhY.gif)

2. Retrieve community member links quickly

![](https://i.imgur.com/k7lH15o.gif)

3. Generate a community leaderboard
4. Generate random hunt modifiers (candidate for change)
5. Generate random weapon loadouts

# How to works
The bot's initial purpose was to replace an obscure randomizer for hunt modifiers (specific animal and weapon loadout to hunt with, etc). That purpose has been fulfilled with the following commands:

1. `s!weapon <reserve> <inventory capacity>`: generates a weapon loadout consisting of 2 or 3 weapons depending on the inventory capacity. Each of the generated weapons can shoot at least 1 animal on the provided reserve. Example usage:

![](https://imgur.com/IpEw1Fz.gif)

2. `s!animal <(optional) reserve>`: generates the name of specific animal to hunt. Optionally, you can provide a reserve which the algorithm will pick an animal from.

![](https://imgur.com/rdfcf0f.gif)

3. `s!reserve`: generates a random reserve to hunt on. Consequently, you can use `s!weapon` or `s!reserve` with the outcome of this command.

# What data is stored and how?
HuntStat does not store any vulnerable data, at all. This is unlikely to change. [Click me to view the related source code](https://github.com/Acygol/huntstat/blob/9c862c1276c98a2574fa147abf7750b0b681c939/framework/database.go#L51-L71)

# Roadmap
1. **Automate updating of TheHunter data:**  
[ammo.json](https://github.com/Acygol/huntstat/blob/master/data/json/ammo.json), [weapons.json](https://github.com/Acygol/huntstat/blob/master/data/json/weapons.json), [reserves.json](https://github.com/Acygol/huntstat/blob/master/data/json/reserves.json), [animals.json](https://github.com/Acygol/huntstat/blob/master/data/json/animals.json) are JSON files containing static data related to the game such as: all existing animals (+ their reserves and their permitted ammo), all existing weapons (+ their ammo), and all existing reserves. If the Expansive Worlds developers make changes to either of these, I must reflect those changes manually.

2. **Command flags to influence randomness** ([issue #11](https://github.com/Acygol/huntstat/issues/11))  
`s!weapon` is currently able to generate odd combinations, such as: 2 bows, 2 muzzleloaders, etc. Using command flags will prevent it from doing so when the user wishes to.

3. **Automatic community member deletion**  
The bot ought to automatically remove community members when they leave your discord server. This functionality is high priority.

4. **Per-community themes and modifiers**  
[...] To be described

5. **Anti-spam measures**  
I agreed to the Discord ToS when I created this bot. One of such is that HuntStat shall not abuse their API. Having anti-spam measures will prevent this from ever happening. This functionality is high priority.

6. **Generate community leaderboards for previous seasons**  
The bot generates a community leaderboard by scraping the widget page of all members. However, the widget only displays data related to the current season, and so there is currently no way to achieve this.

# Build
## libraries:
- [discordgo](https://github.com/bwmarrin/discordgo)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)

## Linux and Windows
- `$ cd /path/to/project/root/`
- `$ go build -o huntstat main/main.go`

When building on windows, you must have a GCC compiler installed (i.e., [mingw_w64](https://mingw-w64.org))

# Run HuntStat for the first time
- Copy the folder `data/json` to your project directory
- Create a folder and call it `config`
- In the config folder, create two `.json` files: `config.json` and `database.json`
- Add the following to `config.json`:
```JSON
{
    "token": "YOUR_BOT_TOKEN",
    "prefix": "s!"
}
```
Change `YOUR_BOT_TOKEN` with your bot toke which you can retrieve from the Discord's developers portal.
- Add the following to `database.json`:
```JSON
{
    "database_file": "./<DATABASE_FILE_NAME>.db"
}
```
Replace `DATABASE_FILE_NAME` with your desired file name.

## Linux
```bash
$ cd /path/to/project/root/
$ chmod +x ./huntstat
$ ./huntstat -init
```

## Windows
```powershell
$ cd \path\to\project\root\
$ .\huntstat.exe -init
```

The `-init` flag populates the database of the necessary tables.
