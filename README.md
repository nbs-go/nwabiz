# WhatsApp for Business Golang Client

`nwabiz` is a HTTP REST Client wrapper for **WhatsApp for Business**. Written for Golang

## Example

> TODO: Complete example

## Testing

### Input

| Environment Key                                    | Description                                            | Required | Value                                           |
| -------------------------------------------------- | ------------------------------------------------------ | -------- | ----------------------------------------------- |
| `NWABIZ_TEST_BASE_URL`                             | WhatsApp for Business Base URL                         | ✓        | URL e.g. `https://111.222.333.444:8080/v1`      |
| `NWABIZ_TEST_USERNAME`                             | WhatsApp for Business Username                         | ✓        | -                                               |
| `NWABIZ_TEST_PASSWORD`                             | WhatsApp for Business Password                         | ✓        | -                                               |
| `NWABIZ_CLIENT_TIMEOUT`                            | HTTP Client Timeout in ms                              | -        | Default value is 10 sec or 10000 ms             |
| `NWABIZ_TEST_CASE_CONTACT_INVALID`                 | Input for Invalid Contact Check                        | ✓        | Use International Notation e.g. `+628123456789` |
| `NWABIZ_TEST_CASE_CONTACT_VALID`                   | Input for Valid Contact Check                          | ✓        | Use International Notation e.g. `+628123456789` |
| `NWABIZ_TEST_INSECURE_SSL`                         | Skip insecure SSL check (e.g. Self-Signed certificate) | -        | Default value is `false`                        |
| `NWABIZ_TEST_CASE_SEND_MESSAGE_TEMPLATE_TEXT_JSON` | Test case for Send Message Template Text               | ✓        | Compact, JSON Formatted. See example below      |

#### `NWABIZ_TEST_CASE_SEND_MESSAGE_TEMPLATE_TEXT_JSON` Example

> - `amount_1000` means the last 3 digits is float precision. For example 8.25 will be 8250

```json
{"to":"<WA_ID>","hsm":{"namespace":"<NAMESPACE>","element_name":"<ELEMENT_NAME>","language":{"code":"<LANG_CODE>"},"localizable_params":[{"default":"<PARAM_VALUE>"},{"default":"99999.99","currency":{"currency_code":"IDR","amount_1000":99999990}}]}}
```                             

> TODO: Complete testing guide

## License

MIT

## Contributors

- Saggaf Arsyad <saggaf@nbs.co.id>
