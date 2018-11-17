[
  {
    "input": "{0..10}",
    "tokens": [
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "EOF",
    ],
    "ast": |||
      (range_numeric [1:6]
        (text [1:2])
        (text [4:6])
        (nil [0:0])
      )
    |||,
    "output": [
      {
        "pattern": "(0|1|2|3|4|5|6|7|8|9|10)",
        "strings": [
          "0",
          "1",
          "2",
          "3",
          "4",
          "5",
          "6",
          "7",
          "8",
          "9",
          "10",
        ],
      },
    ],
  },
  {
    "input": "{0..10..2}",
    "tokens": [
      "BRACE_OPEN",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "RANGE_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "EOF",
    ],
    "ast": |||
      (range_numeric [1:9]
        (text [1:2])
        (text [4:6])
        (text [8:9])
      )
    |||,
    "output": [
      {
        "pattern": "(0|2|4|6|8|10)",
        "strings": [
          "0",
          "2",
          "4",
          "6",
          "8",
          "10",
        ],
      },
    ],
  },
]
