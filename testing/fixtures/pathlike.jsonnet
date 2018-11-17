[
  {
    "input": "{images,thumbs}/*/{0..10}.jpg",
    "tokens": [
      "BRACE_OPEN",
      "LITERAL",
      "LIST_SEPARATOR",
      "LITERAL",
      "BRACE_CLOSE",
      "LITERAL",
      "WILDCARD",
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
      (sequence [1:29]
        (list [1:14]
          (text [1:7])
          (text [8:14])
        )
        (text [15:16])
        (wildcard [16:17])
        (text [17:18])
        (range_numeric [19:24]
          (text [19:20])
          (text [22:24])
          (nil [0:0])
        )
        (text [25:29])
      )
    |||,
    "output": [
      {
        "pattern": "(images|thumbs)/.*?/(0|1|2|3|4|5|6|7|8|9|10)\\.jpg",
        "strings": [
          "images/*/0.jpg",
          "images/*/1.jpg",
          "images/*/2.jpg",
          "images/*/3.jpg",
          "images/*/4.jpg",
          "images/*/5.jpg",
          "images/*/6.jpg",
          "images/*/7.jpg",
          "images/*/8.jpg",
          "images/*/9.jpg",
          "images/*/10.jpg",
          "thumbs/*/0.jpg",
          "thumbs/*/1.jpg",
          "thumbs/*/2.jpg",
          "thumbs/*/3.jpg",
          "thumbs/*/4.jpg",
          "thumbs/*/5.jpg",
          "thumbs/*/6.jpg",
          "thumbs/*/7.jpg",
          "thumbs/*/8.jpg",
          "thumbs/*/9.jpg",
          "thumbs/*/10.jpg",
        ],
      },
    ],
  },
]
