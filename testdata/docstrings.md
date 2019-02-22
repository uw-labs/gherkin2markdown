# DocString variations

## minimalistic

_Given_ a simple DocString

```
first line (no indent)
  second line (indented with two spaces)
third line was empty
```

_Given_ a DocString with content type

```xml
<foo>
  <bar />
</foo>
```

_And_ a DocString with wrong indentation

```
wrongly indented line
```

_And_ a DocString with alternative separator

```
first line
second line
```

_And_ a DocString with normal separator inside

```
first line
"""
third line
```

_And_ a DocString with alternative separator inside

``````
first line
`````
third line
``````

_And_ a DocString with escaped separator inside

```
first line
"""
third line
```
