# Put ideas, plans, todo lists, etc. in here


## Outlines?
A browser "MMO"... where to start. Some ideas on certain game component implementations.

We'll need to have player data in memory to track things like:
- Health
- Inventory

in go something like this

```go
type Player struct {
    Name string
    Health int
    Backpack []string
    CurrentWeapon string
}
```

## Map structure

The meat of the game would be the map/playable zones/world. For a top down game we could manage it through some kind of 3d array or something similar.
For an extremely simple game it could look like this

```
[
[0 0 0 0]
[0 0 0 0]
[0 0 0 0]
[0 3 0 0]
]
```

the different numbers in each location mean different things, but like this it would be hard to track multiple people in one square. It would also be hard to keep track of what the original tile was before a player moved on top of it. So we could do something like this


```
[
[(0, 0, []) (0, 0, [])]
[(0, 0, []) (0, 0, [])]
[(0, 0, []) (0, 0, [])]
[(0, 0, []) (0, 0, [])]
]
```
Each tuple represents a single game tile, where tuple[0] represents the ground type, tuple[1] something else, tuple[2] a list of all players on that location. I think something like this could be a good way to structure our data for managing our game world and keeping track of everything

