# ipfs revival

## Brute Forcing CIDS?

We want to be able to recreate the CID generated from the default IPFS settings by programatically looping though various param settings.

```bash
ipfs add file.json 

# added QmbcdNWDcA9kGpjtVzbFii48E9fonjySxhzbJkzJqjA5Es file.json
# 21 B / 21 B [=========================================] 100.00%
```



```bash
go run main.go 

# THE DATA
# [123 10 32 32 32 32 34 115 116 97 116 117 115 34 58 34 111 107 34 10 125] # 

# THE HASHES
# id              z38cDd
# sha1            z8bvj1ibSdNi7onGpXJu1vWym5qateZgU
# sha2-256        QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n
# sha2-512        zBunRGeXg5ebpZWgv9rCMnqLsD4MacdA1Eycb2ZrzgVdrp8GwhZST6ghFrif98pTBbqRDY6fKbDMzLDCV3TKM9cdApAbX
# sha3-224        z6cYNcWTzKc8PyiWGmWFGovQhbRxqbUedey9HTBb544n
# sha3-256        zdjCksr2tzQHQSy1j5ZyCUWaqALLNjnmpo2qc4HH2nELyk2df
# sha3-384        zM22EdoRk16azRE6pykGArgtjD7dkYSxs6sSC8xzChLE3UeNkU9koZ2CFFuK7r8Fyibq7z3
# sha3-512        zBunor72xa4ncUWgV589Xijz221pXSsWjP5LaANNrnk39QZQ8URiW1JY5vrJYbGeUbobBVe39CT9o7f569zHV7aV2GTZP
# murmur3         zEx8mjdnHy9
# keccak-224      z6cYyW7QDrvugEiPLncbSkX8fg5jb6rueP7DmBP7PZd6
# keccak-256      zdjKK9jsHavQziLBfqo1E6bfa4zvwodpTgaUi2txciJ8nkDh9
# keccak-384      zM27Eo8Y6J1hz4Hq5Uxs9zWRgSjN5gjgf1KPfkkGaeYb4GHmmLnnTxrzMNW561haf4CokrS
# keccak-512      zBurK7Y9DsTrWWA61H24SqjLFTR7EuBQ3LJnMRkSc2T5WQZqZXPAQDdD5d773S6caFwqfxsRcbmuaQrw2DPL2Ac7oBctd
# shake-128       zdjFNvgYFtak7BDHqKkXREuiJzP2P3Vpv66NczQsb4s6gLXe5
# shake-256       zBupkmZT8LNX2doQArJJjqaUcQqkZWEVdgPTRuJXVr639rA5qzB4THZVmFnw9ccnEg7NKUuCMf64ttV5NzHvyRev3btHb
```