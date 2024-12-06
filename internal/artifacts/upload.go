/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package artifacts

import "net/http"

func Upload(a Artifact) error {
	client := http.Client{}
	_ = client
	return nil
}

/*
curl -v -u admin:admin123
-F "maven2.generate-pom=false"
-F "maven2.asset1=@/absolute/path/to/the/local/file/pom.xml"
-F "maven2.asset1.extension=pom"
-F "maven2.asset2=@/absolute/path/to/the/local/file/product-1.0.0.jar;type=application/java-archive"
-F "maven2.asset2.extension=jar"
"http://localhost:8081/service/rest/v1/components?repository=maven-releases"

*/
