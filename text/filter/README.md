## filter

    > cat data/lorem-ipsum-* | filter
      1: Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium
      2: doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore
      3: veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim
      4: ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia
      5: consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque
      6: porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur,
      7: adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore
      8: et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis
      9: nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid
     10: ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea
     11: voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem
     12: eum fugiat quo voluptas nulla pariatur
    >


    > filter data/lorem-ipsum-1 data/lorem-ipsum-2
    ---  data/lorem-ipsum-1
      1: Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium
      2: doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore
      3: veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim
      4: ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia
      5: consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque
      6: porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur,
    ---  data/lorem-ipsum-2
      1: adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore
      2: et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis
      3: nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid
      4: ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea
      5: voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem
      6: eum fugiat quo voluptas nulla pariatur
