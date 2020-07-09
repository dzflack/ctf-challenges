package bid.pleasebecheap.jarjarbinks;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.web.servlet.support.SpringBootServletInitializer;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.context.annotation.Configuration;

@SpringBootApplication
public class JarjarbinksApplication extends SpringBootServletInitializer {

	@Override
    protected SpringApplicationBuilder configure(SpringApplicationBuilder application) {
        return application.sources(JarjarbinksApplication.class);
    }

	public static void main(String[] args) {
		SpringApplication.run(JarjarbinksApplication.class, args);
	}

}
