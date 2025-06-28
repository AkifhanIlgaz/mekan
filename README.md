# Mekan

A CLI application that helps you randomly decide where to go when you're feeling indecisive.

## Installation

First, clone the project:

```bash
git clone https://github.com/AkifhanIlgaz/mekan.git
```

Next, navigate to the project directory and install the application:

```bash
cd mekan
go install
```

You can now use the `mekan` command from anywhere on your system.

## Usage

The primary purpose of the application is to select a random place from a list you create.

### Help Menu

To see all available commands and their descriptions, use:

```bash
mekan -h
```

### Adding a Place

Use the `add` command to add a new place. When you run the command, it will prompt you for the place's name and type (e.g., "cafe", "restaurant", "bar").

```bash
mekan add
```
```
> Name: Starbucks
> Type: Cafe
```

### Listing Places

To view all your saved places in a table, use the `-l` or `--list` flag.

```bash
mekan -l
```

### Selecting a Random Place

To have the application choose a place for you, use the `select` command.

```bash
mekan select
```

If you want to select from a specific type of place, you can use the `-t` or `--type` flag. (The default type is "food").

```bash
mekan select -t cafe
```

After a place is selected, you will be asked to confirm if you want to go there. If you respond with 'y' (yes), the "last visit date" for that place will be updated.

### Deleting a Place

To remove a place from the list, use the `delete` command followed by the place's ID. You can find the IDs by running `mekan -l`.

```bash
mekan delete 1
```

You can also delete multiple places at once:

```bash
mekan delete 1 3 5
```

## Screenshot

![Help Menu](https://raw.githubusercontent.com/AkifhanIlgaz/mekan/main/img/mekan.JPG)
