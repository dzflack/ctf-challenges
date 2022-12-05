# Solution

- Authenticate with default monit creds `admin:monit`
- View post-auth message
- Find [public PoC](https://github.com/dzflack/exploits/blob/master/unix/monit_buffer_overread.py) for vulnerable Monit version (5.25.2)
- Modify PoC to use a custom user agent:

```bash
sed -i "s/'securitytoken=a'/'securitytoken=a', 'User-Agent': 'a'/g" monit_buffer_overread.py
```

- Run PoC and retrieve basic auth credentials
- Use basic auth credentials to access admin view and retrieve flag
