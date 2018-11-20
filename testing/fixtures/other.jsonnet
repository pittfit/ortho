[
  {
    "input": "a,b",
    "tokens": [
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (text [0:3])
    |||,
    "output": [
      {
        "pattern": "a,b",
        "strings": [
          "a,b",
        ],
      },
    ],
  },
  {
    "input": "x{a,b}y",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:7]
        (text [0:1])
        (list [2:5]
          (text [2:3])
          (text [4:5])
        )
        (text [6:7])
      )
    |||,
    "output": [
      {
        "pattern": "x(a|b)y",
        "strings": [
          "xay",
          "xby",
        ],
      },
    ],
  },
  {
    "input": "x{a,{b,c}}y",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:11]
        (text [0:1])
        (list [2:8]
          (text [2:3])
          (list [5:8]
            (text [5:6])
            (text [7:8])
          )
        )
        (text [10:11])
      )
    |||,
    "output": [
      {
        "pattern": "x(a|(b|c))y",
        "strings": [
          "xay",
          "xby",
          "xcy",
        ],
      },
    ],
  },
  {
    "input": "x{1..2}y",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:8]
        (text [0:1])
        (range_numeric [2:6]
          (text [2:3])
          (text [5:6])
          (nil [0:0])
        )
        (text [7:8])
      )
    |||,
    "output": [
      {
        "pattern": "x(1|2)y",
        "strings": [
          "x1y",
          "x2y",
        ],
      },
    ],
  },
  {
    "input": "{a,{1..2}}y",
    "tokens": [
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [1:11]
        (list [1:8]
          (text [1:2])
          (range_numeric [4:8]
            (text [4:5])
            (text [7:8])
            (nil [0:0])
          )
        )
        (text [10:11])
      )
    |||,
    "output": [
      {
        "pattern": "(a|(1|2))y",
        "strings": [
          "ay",
          "1y",
          "2y",
        ],
      },
    ],
  },
  {
    "input": "ab{c,{d,{0..6..2}}}",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "BRACE_CLOSE",
      "BRACE_CLOSE",
      "EOF",
    ],
    "ast": |||
      (sequence [0:16]
        (text [0:2])
        (list [3:16]
          (text [3:4])
          (list [6:16]
            (text [6:7])
            (range_numeric [9:16]
              (text [9:10])
              (text [12:13])
              (text [15:16])
            )
          )
        )
      )
    |||,
    "output": [
      {
        "pattern": "ab(c|(d|(0|2|4|6)))",
        "strings": [
          "abc",
          "abd",
          "ab0",
          "ab2",
          "ab4",
          "ab6",
        ],
      },
    ],
  },
  {
    "input": "x{52,{55..60}}y",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:15]
        (text [0:1])
        (list [2:12]
          (text [2:4])
          (range_numeric [6:12]
            (text [6:8])
            (text [10:12])
            (nil [0:0])
          )
        )
        (text [14:15])
      )
    |||,
    "output": [
      {
        "pattern": "x(52|(55|56|57|58|59|60))y",
        "strings": [
          "x52y",
          "x55y",
          "x56y",
          "x57y",
          "x58y",
          "x59y",
          "x60y",
        ],
      },
    ],
  },
  {
    "input": "x{52,y{55..60},}z",
    "tokens": [
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LIST_SEPARATOR",
      "BRACE_CLOSE",
      "LITERAL",
      "EOF",
    ],
    "ast": |||
      (sequence [0:17]
        (text [0:1])
        (list [0:13]
          (text [2:4])
          (sequence [5:13]
            (text [5:6])
            (range_numeric [7:13]
              (text [7:9])
              (text [11:13])
              (nil [0:0])
            )
          )
          (nil [0:0])
        )
        (text [16:17])
      )
    |||,
    "output": [
      {
        "pattern": "x(52|y(55|56|57|58|59|60)|)z",
        "strings": [
          "x52z",
          "xy55z",
          "xy56z",
          "xy57z",
          "xy58z",
          "xy59z",
          "xy60z",
          "xz",
        ],
      },
    ],
  },
]
