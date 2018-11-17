[
  {
    "input": "",
    "tokens": [
      "EOF",
    ],
    "ast": |||
      (sequence [0:0])
    |||,
    "output": [
      {
        "pattern": "",
        "strings": [],
      },
    ],
  },
  {
    "input": "{,}",
    "tokens": [
      "BRACE_OPEN",
      "LIST_SEPARATOR",
      "BRACE_CLOSE",
      "EOF",
    ],
    "ast": |||
      (list [0:0]
        (nil [0:0])
        (nil [0:0])
      )
    |||,
    "output": [
      {
        "pattern": "(|)",
        "strings": ["", ""],
      },
    ],
  },
  {
    "input": "a{2,1}b{X,Y,X}c",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:15]
        (text [0:1])
        (list [2:5]
          (text [2:3])
          (text [4:5])
        )
        (text [6:7])
        (list [8:13]
          (text [8:9])
          (text [10:11])
          (text [12:13])
        )
        (text [14:15])
      )
    |||,
    "output": [
      {
        "pattern": "a(2|1)b(X|Y|X)c",
        "strings": [
          "a2bXc",
          "a2bYc",
          "a2bXc",
          "a1bXc",
          "a1bYc",
          "a1bXc",
        ],
      },
    ],
  },
]
