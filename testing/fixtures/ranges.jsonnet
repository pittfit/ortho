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
    "input": "{10..0}",
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
        (text [1:3])
        (text [5:6])
        (nil [0:0])
      )
    |||,
    "output": [
      {
        "pattern": "(10|9|8|7|6|5|4|3|2|1|0)",
        "strings": [
          "10",
          "9",
          "8",
          "7",
          "6",
          "5",
          "4",
          "3",
          "2",
          "1",
          "0",
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
  {
    "input": "{0..10..-2}",
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
      (range_numeric [1:10]
        (text [1:2])
        (text [4:6])
        (text [8:10])
      )
    |||,
    "output": [
      {
        "pattern": "(10|8|6|4|2|0)",
        "strings": [
          "10",
          "8",
          "6",
          "4",
          "2",
          "0",
        ],
      },
    ],
  },
  {
    "input": "{10..0..2}",
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
        (text [1:3])
        (text [5:6])
        (text [8:9])
      )
    |||,
    "output": [
      {
        "pattern": "(10|8|6|4|2|0)",
        "strings": [
          "10",
          "8",
          "6",
          "4",
          "2",
          "0",
        ],
      },
    ],
  },
  {
    "input": "{10..0..-2}",
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
      (range_numeric [1:10]
        (text [1:3])
        (text [5:6])
        (text [8:10])
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
