01   PKLR1-DETAIL-LOAN-RECORD.                                   
    10  PKLR1-BASIC-SECTION.                                     
        20  PKLR1-SORT-CONTROL-FIELD.                            
            30  PKLR1-USER-IDENT         PIC X(1).               
            30  PKLR1-EXTRACT-CODE       PIC X(1).               
                88  PKLR1-DATE-RECORD            VALUE '*'.      
                88  PKLR1-DATA-RECORD            VALUE '0'.      
                88  PKLR1-END-OF-FILE            VALUE '9'.      
            30  PKLR1-SECTION            PIC X(1).               
            30  PKLR1-TYPE               PIC X(1).               
            30  PKLR1-NUMERIC-STATE-CODE PIC X(2).               
            30  PKLR1-CONTRACT-NUMBER    PIC X(10).              
        20  PKLR1-PAR-PEN-REG-CODE       PIC X(1).               
        20  PKLR1-VALUATION-CODE.                                
            30  PKLR1-MORTALITY-TABLE    PIC X(2).               
            30  PKLR1-LIVES-CODE         PIC X(1).               
            30  PKLR1-FUNCTION           PIC X(1).               
            30  PKLR1-VAL-INTEREST       PIC S9(2)V9(3).  
            30  PKLR1-MODIFICATION       PIC X(1).               
            30  PKLR1-INSURANCE-CLASS    PIC X(1).               
            30  PKLR1-SERIES             PIC X(5).               
        20  PKLR1-POLICY-STATUS          PIC X(2).               
        20  PKLR1-PAR-CODES.                                     
            30  PKLR1-PAR-TYPE           PIC X(1).               
            30  PKLR1-DIVIDEND-OPTION    PIC X(1).               
            30  PKLR1-OTHER-OPTION       PIC X(1).               
        20  PKLR1-ALPHA-STATE-CODE       PIC X(2).  
        20  PKLR1-ALPHA-STATE-CODEX1       PIC XX.
        20  PKLR1-ALPHA-STATE-CODEX2       PIC 99.
        20  PKLR1-ALPHA-STATE-CODEX3       PIC 9(3).
        20  PKLR1-ALPHA-STATE-CODEX4       PIC 9(3)V99.  