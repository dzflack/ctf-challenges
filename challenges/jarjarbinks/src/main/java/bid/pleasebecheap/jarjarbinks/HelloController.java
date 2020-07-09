package bid.pleasebecheap.jarjarbinks;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import java.io.IOException;
import java.net.JarURLConnection;
import java.net.URL;
import java.net.SocketTimeoutException;
import java.util.Enumeration;
import java.util.jar.JarEntry;
import java.util.jar.JarFile;
import java.util.ArrayList;
import java.util.List;
import org.springframework.web.bind.annotation.RequestMethod;

@RestController
public class HelloController {

    @RequestMapping(value = "/rest/jar", method = {RequestMethod.POST})
    public UnpackResult rest(@RequestBody RemoteFile remoteFile) {
        try {
            List<String> contents = downloadUnzipFile(remoteFile.getFileUrl());
            return new UnpackResult(contents);
        } catch (Exception e) {
            return new UnpackResult(e);
        }
    }

    public List<String> downloadUnzipFile(final String fileUrl) throws IOException, SocketTimeoutException {
        final URL url = new URL(fileUrl);
        
        JarURLConnection connection = (JarURLConnection) url.openConnection();
        connection.setConnectTimeout(4000);
        connection.setReadTimeout(4000);

        System.out.print("**********  Content Length is: " + connection.getContentLength() + "    **********\n");
        if (connection.getContentLength() > 1000000 || connection.getContentLength() < 0){
            throw new SocketTimeoutException();
        }

        JarFile jar = connection.getJarFile();
        Enumeration<JarEntry> enumEntries = jar.entries();
        List<String> contents = new ArrayList<>();
        
        while (enumEntries.hasMoreElements()) {
            java.util.jar.JarEntry file = (java.util.jar.JarEntry) enumEntries.nextElement();
            String destDir = "/tmp/";
			java.io.File f = new java.io.File(destDir  + java.io.File.separator + file.getName());
            if (file.isDirectory()) { 
                f.mkdir();
                continue;
            }
            java.io.InputStream is = jar.getInputStream(file);
            java.io.FileOutputStream fos = new java.io.FileOutputStream(f);
            while (is.available() > 0) {
                fos.write(is.read());
            }
            fos.close();
            is.close();
            contents.add(file.getName());
        }
        jar.close();

        return contents;
    }
}