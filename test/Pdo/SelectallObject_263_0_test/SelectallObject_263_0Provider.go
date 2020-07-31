package SelectallObject_263_0_test

type  SelectallObject_263_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            `[{"A1":"a1","Dd":"hahaa","Id":10},{"A1":"a1","Dd":"hahaa","Id":11}]`,
            map[string]interface{}{"sql":"select id,a1,dd from  a where id in (?,?)","binds":[]interface{}{10,11}},
            " - id=127",
        },
	}

}