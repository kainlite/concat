### Readme for concat
Concatenate stuff with a small yaml config

Example config:
```yaml
outdir: out
indir: example
chunks:
  id_a:
    filename: "fancy_name_a.txt"
    parts:
      - base
      - body_a
      - foot
  id_b:
    filename: "fancy_name_b.txt"
    parts:
      - base
      - body_b
      - foot
  id_c:
    filename: "fancy_name_c.txt"
    parts:
      - base
      - body_c
      - foot
```

Outcome 3 files, with different body with same base and foot, be advised this is more of a toy project than a real tool.

Questions like: why would you do that instead of using bash and cat?
Short answer: Where is the fun in that? I'm playing with go :)

### TODO:
- Refactor
- Add tests
