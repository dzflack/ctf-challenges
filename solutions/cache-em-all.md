# Solution

1. Download [dnspy](https://github.com/dnSpy/dnSpy/releases)
2. Open crypto-3.exe and crypto-3-library.dll in dnspy
3. Replace the `decrypt` function in `crypto-3.crypto_3.Form1` with the function below
4. At the end of the `button1_Click` function in `crypto-3.crypto_3.Form1` add the decrypt call, as shown below
5. Save the binary to disk
6. Run crypto-3.exe, supply random inputs to username and password, observe flag in message box

Decrypt function

```C#
private void decrypt(X509Certificate2 certificate, string currentDirectory)
{
    var privateKey = certificate.PrivateKey as RSACryptoServiceProvider;
    string path = currentDirectory + "\\.cachedCredentials";
    using (StreamReader sr = new StreamReader(path))
    {
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            var encryptedBytes = System.Convert.FromBase64String(line);
            var data = privateKey.Decrypt(encryptedBytes, false);
            MessageBox.Show(System.Text.Encoding.UTF8.GetString(data), "decrypted", MessageBoxButtons.OK, MessageBoxIcon.Information);
        }
    }
}

Decrypt Call

```C#
this.decrypt(x509Certificate2Collection[0], currentDirectory);
```
