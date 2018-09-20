### Readme for concat
Concatenate stuff with a small yaml config

Example config:
```yaml
outdir: out
chunks: example
values:
  file_a:
    base
    body_a
    foot
  file_b:
    base
    body_b
    foot
  file_c:
    base
    body_c
    foot
```

Outcome 3 files, with different body with same base and foot, be advised this is more of a toy project than a real tool.

### TODO:
- Refactor
- Add tests
