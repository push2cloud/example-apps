package com.swisscom.cloud.cloudfoundry.sampleapp.java;
 
import static spark.Spark.get;
import static spark.Spark.post;
import static spark.Spark.port;

import java.io.IOException;
import java.io.StringWriter;
import java.math.BigDecimal;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;

public class ProductService {

    private static final int HTTP_OK = 200;
    private static final int HTTP_BAD_REQUEST = 400;
    
	private static final Info INFO = new Info("Your Java sample app is up and running !", "1.0.00", System.getenv().get("APP_MODE"));
    
    public static void main( String[] args) {
    	
    	port(getCloudAssignedPort());
    	
    	ProductRepository productRepository = getProductRepository();
    	
        get("/", (request, response) -> {
            response.status(HTTP_OK);
            response.type("application/json");
            return mapToJson(INFO);
        });
        
        post("/products", (request, response) -> {
            try {
                ObjectMapper mapper = new ObjectMapper();
                Product product = mapper.readValue(request.body(), Product.class);
                if (! product.isValid()) {
                    response.status(HTTP_BAD_REQUEST);
                    return "Product input invalid";
                }
                long id = productRepository.add(product);
                response.status(HTTP_OK);
                response.type("application/json");
                return id;
            } catch (JsonParseException exception) {
                response.status(HTTP_BAD_REQUEST);
                return "Json payload invalid";
            }
        });
        
        get("/products", (request, response) -> {
            response.status(HTTP_OK);
            response.type("application/json");
            return mapToJson(productRepository.findAll());
        });
    }

    private static String mapToJson(Object data) {
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.enable(SerializationFeature.INDENT_OUTPUT);
            StringWriter writer = new StringWriter();
            mapper.writeValue(writer, data);
            return writer.toString();
        } catch (IOException exception){
            throw new RuntimeException("Could not write JSON response payload", exception);
        }
    }
    
    private static int getCloudAssignedPort() {
        ProcessBuilder processBuilder = new ProcessBuilder();
        if (processBuilder.environment().get("PORT") != null) {
            return Integer.parseInt(processBuilder.environment().get("PORT"));
        }
        return 4567; //return default port if cloud port isn't set (i.e. on localhost)
    }
    
    private static ProductRepository getProductRepository() {    	
       	return new SimpleProductRepository();
    }
    
    public static class Info {
      	 
    	private String status;
        private String version;
        private String appMode;
        
        public Info(String status, String version, String appMode) {
        	this.status = status;
        	this.version = version;
        	this.appMode = appMode;
        }
       
        public String getStatus() {
        	return status;
        }
        
        public String getVersion() {
        	return version;
        }
        
        public String getAppMode() {
        	return appMode;
        }
    }
    
    public static class Product {
   	 
    	private long id;
        private String description;
        private BigDecimal price;
       
        public long getId() {
        	return id;
        }
        
        public void setId(long id) {
        	this.id = id;
        }
        
        public String getDescription() {
        	return description;
        }
        
        public BigDecimal getPrice() {
        	return price;
        }
        
        @JsonIgnore
        public boolean isValid() {
        	if (description == null || description.isEmpty()) {
        		return false;
        	}
        	if (price == null) {
        		return false;
        	}
        	return true;
        }
    }
    
    public interface ProductRepository {    	
    	public long add(Product product);   	
    	public Collection<Product> findAll();
    }
    
    public static class SimpleProductRepository implements ProductRepository {
    	
    	private Map<Long,Product> products = new HashMap<Long,Product>();
    	private long id;
    	   	
    	@Override
    	public long add(Product product) {
    		product.setId(nextId());
    		products.put(product.getId(), product);
    		return product.getId();
    	}
    	
    	@Override
    	public Collection<Product> findAll() {
    		return products.values();
    	}
    	
    	private long nextId() {
    		return ++id;
    	}
    }
}