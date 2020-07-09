package bid.pleasebecheap.jarjarbinks;

import com.fasterxml.jackson.annotation.JsonCreator;

public class RemoteFile {
    private final String fileUrl;

    @JsonCreator
    public RemoteFile(String fileUrl) {
        this.fileUrl = fileUrl;
    }

    public String getFileUrl() {
        return fileUrl;
    }
}