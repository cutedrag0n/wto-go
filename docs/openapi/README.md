# Official specification snapshots

Retrieved on 2026-11-04 from the [WTO Developer Portal](https://apiportal.wto.org/apis), using its public API-definition export (OpenAPI 3 JSON).

The portal's `config.json` identifies its public management endpoint:
`https://wto-data.management.azure-api.net/subscriptions/000/resourceGroups/000/providers/Microsoft.ApiManagement/service/WTO-DATA`.

Export route: `/apis/{id}?api-version=2022-04-01-preview&format=openapi%2Bjson&export=true`.
The files preserve each export's `value` object without semantic changes.

| API ID | File | Operations |
|---|---|---|
| version1 | version1.json | 18 |
| eping | eping.json | 2 |
| quantitative-restrictions | quantitative-restrictions.json | 6 |
| tfad | tfad.json | 1 |

These are upstream documentation, including upstream inconsistencies. SDK-specific observations belong in ../api-notes.md.
