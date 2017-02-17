package route53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "domio_api/components/config"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/fatih/color"
    "fmt"
    "time"
    "log"
)

func DeleteDomainZoneReal(domain *domiodb.Domain) error {
    conf := config.Config
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return err
    }

    svc := route53.New(sess)

    params := &route53.DeleteHostedZoneInput{
        Id: &domain.ZoneId.String,
    }
    //resp, err := svc.DeleteHostedZone(params)
    resp, err := svc.DeleteHostedZone(params)

    if err != nil {
        color.Set(color.FgRed)
        log.Println(err)
        color.Unset()
        return err
    }
    log.Println(resp)
    log.Print("Domain zone removed from Route 53")
    return nil
}

func CreateDomainZoneReal(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {
    conf := config.Config
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})
    if err != nil {
        fmt.Println("failed to create session,", err)
        return &route53.CreateHostedZoneOutput{}, err
    }

    r53Service := route53.New(sess)
    id := time.Now().Format(time.RFC850);

    params := &route53.CreateHostedZoneInput{
        CallerReference: &id,
        Name:            aws.String(domain.Name),
    }

    resp, err := r53Service.CreateHostedZone(params)

    if err != nil {
        color.Set(color.FgRed)
        log.Println(params.CallerReference)
        log.Println(id)
        log.Println(err)
        color.Unset()
        return nil, err
    }

    return resp, nil

}
