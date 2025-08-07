# Books

This repo contains a simple TUI to manage a book collection.

![list view](gallery/list.png)
![filer view](gallery/filter.png)
![book view](gallery/book.png)


## Running

The syntax to run this application is

```
go run . <dbPath> <programName/"list"> [additionalParams]
```

or, if you wish

```
go build .
./books <dbPath> <programName/"list"> [additionalParams]
```

- `dbPath` is the path of a sqlite database. Does not need to exist in the
  filesystem.

- `programName` is the name of the program to run.
  + `list` shows the list of available programs
  + `hierarchy` lets you navigate through the collection hierarchy (libraries,
    groups, sagas and books)
  + `flat` lets you navigate the list of all books, without a "folderized" view
  + `book` lets you see directly the details of a book, whose isbn will be
    specified in the `additionalPrams`.

## Further Explanation

### ... I am sorry, what?

Okay, so.

I had this books list, and I divided it into `libraries` (e.g. _Novels_ and
_Comics_), `groups` (e.g. _Fantasy_ and _Sci-Fi_ for the _Novels_ library) and
`sagas` (e.g. _Nevernight_).

With this project, I can search through the collections and view the books data
(with an ASCII thumbnail preview!).

### okay?

Yeah I know this is quite stupid, but I am working on it.

Right now the data is read-only: I'm still trying to build a pleasing interface,
the editing functionalities will come later (if I'll ever get there).

### uhm, yes, but...

I understand that many of the things done (or thinked) in this project may not
fit everyone, or even anyone, but I preferred to start developing something
suitable for my needs, not the most general thing ever.

That said, I hope somebody else will use it some day.

## How did you do it?

The project relies on some [charmbracelet](https://charm.land) libraries, i.e.
[bubbletea](https://github.com/charmbracelet/bubbletea),
[bubbles](https://github.com/charmbracelet/bubbles) and
[lipgloss](https://github.com/charmbracelet/lipgloss).

I then made a `components` folder, containing some reusable, data-agnostic
models (like `list` to see a list and `tabs` for a tabbed view), used by the
`models.go` file, which contains all the specialized models, used by the various
programs, described iside `p_*.go` files.

## Last notes

This project is released under the MIT license. See the `LICENSE` file for more
details.

Any contribution is highly appreciated.
