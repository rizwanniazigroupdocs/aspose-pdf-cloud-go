![](https://img.shields.io/badge/api-v3.0-lightgrey) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/aspose-pdf-cloud/aspose-pdf-cloud-go) [![GitHub license](https://img.shields.io/github/license/aspose-pdf-cloud/aspose-pdf-cloud-go)](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-go)

# Aspose.PDF Cloud SDK for GO
[Aspose.PDF Cloud](https://products.aspose.cloud/pdf) is a true REST API that enables you to perform a wide range of document processing operations including creation, manipulation, conversion and rendering of PDF documents in the cloud.

Our Cloud SDKs are wrappers around REST API in various programming languages, allowing you to process documents in language of your choice quickly and easily, gaining all benefits of strong types and IDE highlights. This repository contains new generation SDKs for Aspose.PDF Cloud and examples.

These SDKs are now fully supported. If you have any questions, see any bugs or have enhancement request, feel free to reach out to us at [Free Support Forums](https://forum.aspose.cloud/c/pdf).

Extract Text & Images of a PDF document online https://products.aspose.app/pdf/parser.
## Key Features

* Work with Aspose for Cloud storage.
* Create PDF documents from different formats.
* Convert local or remote PDF documents to different formats.
* Split one PDF document to pages and save as separate documents.
* Merge multiple PDF documents.
* Get word count of all PDF pages.
* Manipulate document properties.
* Add, copy, move or delete PDF pages.
* Convert all or specific pages to other formats.
* Update form fields.
* Replace Images in a PDF file.
* Get all annotations from a PDF page.
* Get all links from a PDF page.
* Get all attachments from a PDF.

### New Features & Recent API Changes

#### 20.10
* A new version of Aspose.PDF Cloud was prepared using the latest version of Aspose.PDF for .Net.

#### 20.9
* Implemented custom fonts for Annotation API.
* Implemented possibility to apply redaction annotation. Now you can apply annotation when adding or updating.

#### 20.8
* Implemented custom fonts for Text API.
* Added custom fonts for Table API.
* Added support for custom fonts for Stamps API.
* Support for custom fonts for Header/Footer API.
* Included custom fonts for Replace Text API.

#### 20.7
* Added Support for PDF_A_3A Format.
* Included support for "MaxResolution" option in OptimizeOption.
* Implemented Info action.
* Support ImageCompressionOptions in OptimizeOptions.


## Installation
Put the package under your project folder and add the following in import:
```
    "./asposepdfcloud"
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```go
func GetDocumentCircleAnnotations() (CircleAnnotationsResponse, *http.Response, error) {
    pdfAPI := NewPdfApiService("AppSid", "AppKey", "")
	name := "PdfWithAnnotations.pdf"	

	args := map[string]interface{} {
		"folder": "path/to/remote/folder",
	}

	return pdfAPI.GetDocumentCircleAnnotations(name, args)
}
```

## Unit Tests
Aspose PDF Cloud SDK includes a suite of unit tests within the "test" subdirectory. These Unit Tests also serves as examples of how to use the Aspose PDF Cloud SDK.

## Licensing
All Aspose.PDF Cloud SDKs are licensed under [MIT License](LICENSE).

## Aspose.PDF Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift|Perl|Go|
|---|---|---|---|---|---|---|--|--|--|
| [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-dotnet) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-java) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-php) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-python) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-ruby)  | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-nodejs) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-android) | [GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-swift)|[GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-perl) |[GitHub](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-go) |
| [NuGet](https://www.nuget.org/packages/Aspose.pdf-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-pdf-cloud) | [Composer](https://packagist.org/packages/aspose/pdf-sdk-php) | [PIP](https://pypi.org/project/asposepdfcloud/) | [GEM](https://rubygems.org/gems/aspose_pdf_cloud)  | [NPM](https://www.npmjs.com/package/asposepdfcloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-pdf-cloud) | [Cocoapods](https://cocoapods.org/pods/AsposepdfCloud)|[Meta Cpan](https://metacpan.org/release/AsposeSlidesCloud-SlidesApi) | [Go.Dev](https://pkg.go.dev/github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/) | 

[Product Page](https://products.aspose.cloud/pdf/go) | [Documentation](https://docs.aspose.cloud/display/pdfcloud/Home) | [API Reference](https://apireference.aspose.cloud/pdf/) | [Code Samples](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-go) | [Blog](https://blog.aspose.cloud/category/pdf/) | [Free Support](https://forum.aspose.cloud/c/pdf) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
